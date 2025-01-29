import prisma from '~/lib/prisma'
import { isAuthenticatedUser } from '~/server/utils/loadAuthenticatedUserFromSession'

export default defineEventHandler(async (event) => {
    const result = await loadAuthenticatedUserFromSession(event)
    if (!isAuthenticatedUser(result)) {
        return result.err
    }

    const validatedBody = await readValidatedBody(event, getTasksSchemas().createModelSchema(result.user).safeParseAsync)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Validation failed', validatedBody.error)
    }

    const task = await prisma.task.create({
        data: {
            status: validatedBody.data.status,
            message: validatedBody.data.message,
            user: { connect: result.user },
            priority: { connect: { id: validatedBody.data.priority_id } },
            category: { connect: { id: validatedBody.data.category_id } },
        },
    })

    setResponseStatus(event, 201)

    return {
        status: 201,
        data: task,
    }
})
