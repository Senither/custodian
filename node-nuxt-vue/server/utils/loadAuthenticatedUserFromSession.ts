import type { User } from '@prisma/client'
import prisma from '~/lib/prisma'

export type AuthenticatedUserResult =
    | { user: User, err: null }
    | { user: null, err: any }

/**
 * Loads the authenticated user from the database using the available
 * session data, if it fails to load the authenticated user it will
 * instead return an error response object.
 *
 * @param event The event object
 * @returns The authenticated user or an error
 */
export default async function (event: any): Promise<AuthenticatedUserResult> {
    const session = await getUserSession(event)
    if (Object.keys(session).length === 0) {
        return wrapError(createErrorResponse(event, 401, 'Unauthorized'))
    }

    if (!session.user || !('id' in session.user)) {
        return wrapError(createErrorResponse(event, 400, 'Invalid session data'))
    }

    const userId = session.user.id ?? 0
    if (!userId) {
        return wrapError(createErrorResponse(event, 400, 'Invalid session data'))
    }

    const user: User | null = await prisma.user.findFirst({
        where: { id: userId },
    })

    if (!user) {
        return wrapError(createErrorResponse(event, 404, 'Found no user matching session data'))
    }

    return { user, err: null }
}

/**
 * Checks if the given result is an authenticated user.
 *
 * @param result The result that should be checked
 * @returns Whether the result is an authenticated user
 */
export function isAuthenticatedUser(result: AuthenticatedUserResult): result is { user: User, err: null } {
    return result.err === null
}

/**
 * Wraps the given error response in the expected format and returns it.
 *
 * @param err The error that should be returned
 * @returns The error wrapped in an object
 */
function wrapError(err: any): { user: null, err: any } {
    return { user: null, err }
}
