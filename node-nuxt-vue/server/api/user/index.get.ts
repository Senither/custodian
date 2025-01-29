import { isAuthenticatedUser } from '~/server/utils/loadAuthenticatedUserFromSession'

export default defineEventHandler(async (event) => {
    const result = await loadAuthenticatedUserFromSession(event)
    if (!isAuthenticatedUser(result)) {
        return result.err
    }

    return {
        status: 200,
        data: {
            id: result.user.id,
            name: result.user.name,
            email: result.user.email,
            email_verified_at: result.user.email_verified_at,
            created_at: result.user.created_at,
        },
    }
})
