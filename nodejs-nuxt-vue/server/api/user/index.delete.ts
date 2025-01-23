import { z } from 'zod'
import prisma from '~/lib/prisma'
import { isAuthenticatedUser } from '~/server/utils/loadAuthenticatedUserFromSession'

const schema = z.object({
    password: z.string(),
})

export default defineEventHandler(async (event) => {
    const validatedBody = await readValidatedBody(event, schema.safeParse)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Validation error', validatedBody.error)
    }

    const result = await loadAuthenticatedUserFromSession(event)
    if (!isAuthenticatedUser(result)) {
        return result.err
    }

    if (!result.user.password || !await verifyPassword(result.user.password, validatedBody.data.password)) {
        return createErrorResponse(event, 400, 'Invalid password provided')
    }

    await prisma.user.delete({
        where: {
            id: result.user.id,
        },
    })

    setResponseStatus(event, 202)

    return {
        status: 202,
        message: 'Your account has been deleted',
    }
})
