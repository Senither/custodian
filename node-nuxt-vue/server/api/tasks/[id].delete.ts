import prisma from '~/lib/prisma'
import { isAuthenticatedUser } from '~/server/utils/loadAuthenticatedUserFromSession'

export default defineEventHandler(async (event) => {
    const result = await loadAuthenticatedUserFromSession(event)
    if (!isAuthenticatedUser(result)) {
        return result.err
    }

    const params = await getValidatedRouterParams(event, getTasksSchemas().param.safeParse)
    if (!params.success) {
        return createErrorResponse(event, 400, 'Invalid task ID provided', params.error)
    }

    const task = await prisma.task.count({
        where: {
            id: params.data.id,
            user: result.user,
        },
    })

    if (task === 0) {
        return createErrorResponse(event, 404, 'No record found for the provided task ID')
    }

    await prisma.task.delete({
        where: { id: params.data.id },
    })

    setResponseStatus(event, 202)

    return {
        status: 202,
        message: 'The task has been deleted',
    }
})
