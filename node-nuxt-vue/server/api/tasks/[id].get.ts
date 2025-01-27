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

    const task = await prisma.task.findFirst({
        where: {
            id: params.data.id,
            user: result.user,
        },
        include: {
            priority: getTasksSchemas().includeQuery,
            category: getTasksSchemas().includeQuery,
        },
    })

    if (!task) {
        return createErrorResponse(event, 404, 'No record found for the provided task ID')
    }

    return {
        status: 200,
        data: task,
    }
})
