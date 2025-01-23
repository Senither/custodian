
export default defineEventHandler(async (event) => {
    const { user, err } = await loadAuthenticatedUserFromSession(event)
    if (err) {
        return err
    }

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
