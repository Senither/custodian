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
        throw createError({
            statusCode: 400,
            message: 'Missing required fields',
            data: validatedBody.error,
        })
    }

    const user: User | null = await prisma.user.findFirst({
        where: { email: validatedBody.data.email },
    })

    if (!user) {
        throw createError({
            statusCode: 401,
            message: 'Invalid email or password'
        })
    }

    if (!user.password || !await verifyPassword(user.password, validatedBody.data.password)) {
        throw createError({
            statusCode: 401,
            message: 'Invalid email or password'
        })
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
