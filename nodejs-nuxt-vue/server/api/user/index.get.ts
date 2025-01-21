import prisma from '~/lib/prisma'

export default defineEventHandler(async (event) => {
    const session = await getUserSession(event)
    if (Object.keys(session).length === 0) {
        return createErrorResponse(event, 401, 'Unauthorized')
    }

    const userId = session.user?.id ?? 0
    if (!userId) {
        return createErrorResponse(event, 400, 'Invalid session data')
    }

    const user = await prisma.user.findFirstOrThrow({
        where: { id: userId },
    })

    return {
        status: 200,
        data: {
            id: user.id,
            name: user.name,
            email: user.email,
            email_verified_at: user.email_verified_at,
            created_at: user.created_at,
        },
    }
})
