import prisma from '~/lib/prisma'
import { isAuthenticatedUser } from '~/server/utils/loadAuthenticatedUserFromSession'

export default defineEventHandler(async (event) => {
    const result = await loadAuthenticatedUserFromSession(event)
    if (!isAuthenticatedUser(result)) {
        return result.err
    }

    return {
        status: 200,
        data: await prisma.task.findMany({
            where: { user: result.user },
            include: {
                priority: getTasksSchemas().includeQuery,
                category: getTasksSchemas().includeQuery,
            },
        }),
    }
})
