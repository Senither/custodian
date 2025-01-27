import type { User } from '@prisma/client'
import { z } from 'zod'
import prisma from '~/lib/prisma'

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

    createDefaultUserRecords(user).catch(console.error)

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
        },
    }
})
