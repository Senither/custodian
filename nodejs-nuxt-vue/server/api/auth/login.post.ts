import prisma from '~/lib/prisma'
import { z } from 'zod'
import { User } from '@prisma/client'

const schema = z.object({
    email: z.string().email(),
    password: z.string(),
})

export default defineEventHandler(async (event) => {
    const validatedBody = await readValidatedBody(event, schema.safeParse)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Missing required fields', validatedBody.error)
    }

    const user: User | null = await prisma.user.findFirst({
        where: { email: validatedBody.data.email },
    })

    if (!user || !user.password || !await verifyPassword(user.password, validatedBody.data.password)) {
        return createErrorResponse(event, 401, 'Invalid email or password')
    }

    await replaceUserSession(event, {
        user: {
            id: user.id,
            name: user.name,
        },
    })

    return {
        statusCode: 200,
        message: 'Login successful',
        user: {
            id: user.id,
            name: user.name,
        }
    }
})
