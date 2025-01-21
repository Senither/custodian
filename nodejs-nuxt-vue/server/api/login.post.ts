import prisma from '~/lib/prisma'
import { z } from 'zod'
import { User } from '@prisma/client'

const schema = z.object({
    email: z.string().email(),
    password: z.string(),
})

export default defineEventHandler(async (event) => {
    const { email, password } = await readValidatedBody(event, schema.parse)

    const user: User | null = await prisma.user.findFirst({
        where: { email },
    })

    if (!user) {
        throw createError({
            statusCode: 401,
            message: 'Bad credentials'
        })
    }

    if (!user.password || !await verifyPassword(user.password, password)) {
        throw createError({
            statusCode: 401,
            message: 'Bad credentials'
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
