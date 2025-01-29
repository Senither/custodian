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

    const validatedBody = await readValidatedBody(event, getTasksSchemas().updateModelSchema(result.user).safeParseAsync)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Validation failed', validatedBody.error)
    }

    const updateBody: any = {
        status: validatedBody.data.status,
        message: validatedBody.data.message,
    }

    if (validatedBody.data.priority_id) {
        updateBody.priority = { connect: { id: validatedBody.data.priority_id } }
    }

    if (validatedBody.data.category_id) {
        updateBody.category = { connect: { id: validatedBody.data.category_id } }
    }

    const task = await prisma.task.update({
        where: {
            id: params.data.id,
            user: result.user,
        },
        data: updateBody,
    })

    return {
        status: 200,
        message: 'Task details updated successfully',
        data: task,
    }
})
