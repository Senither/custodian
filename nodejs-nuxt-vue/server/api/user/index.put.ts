import { z } from 'zod'
import prisma from '~/lib/prisma'

export default defineEventHandler(async (event) => {
    const { user, err } = await loadAuthenticatedUserFromSession(event)
    if (err) {
        return err
    }

    const validatedBody = await readValidatedBody(event, z.object({
        name: z.string().min(2).optional(),
        email: z.string().email().optional(),
    }).refine(async data => {
        if (!data.email) {
            return true
        }
        return ! await checkUserWithEmailAlreadyExists(data.email, user.id)
    }, {
        message: 'Email has already been taken',
        path: ['email'],
    }).safeParseAsync)

    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Missing required fields', validatedBody.error)
    }

    await prisma.user.update({
        where: { id: user.id },
        data: validatedBody.data,
    })

    return {
        status: 200,
        data: validatedBody
    }
})
