import prisma from '~/lib/prisma'
import { z } from 'zod'
import { User } from '@prisma/client'

const schema = z.object({
    name: z.string().min(2),
    email: z.string().email(),
    password: z.string().min(8),
    password_confirmation: z.string(),
}).refine(data => data.password === data.password_confirmation, {
    message: 'Password and password confirmation do not match',
    path: ['password'],
}).refine(async data => !await checkUserWithEmailAlreadyExists(data.email), {
    message: 'Email has already been taken',
    path: ['email'],
})

export default defineEventHandler(async (event) => {
    const validatedBody = await readValidatedBody(event, schema.safeParseAsync)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Missing required fields', validatedBody.error)
    }

    const hashedPassword = await hashPassword(validatedBody.data.password)

    const user: User = await prisma.user.create({
        data: {
            name: validatedBody.data.name,
            email: validatedBody.data.email,
            password: hashedPassword,
        },
    })

    await replaceUserSession(event, {
        user: {
            id: user.id,
            name: user.name,
        },
    })

    setResponseStatus(event, 201)

    return {
        statusCode: 201,
        message: 'Register successful',
        user: {
            id: user.id,
            name: user.name,
        }
    }
})

/**
 * Checks if a user with the given email already exists.
 *
 * @param email The email to check
 * @returns Promise<boolean> Whether a user with the email already exists
 */
async function checkUserWithEmailAlreadyExists(email: string): Promise<boolean> {
    const user = await prisma.user.findFirst({
        where: { email },
    })

    return user !== null
}
