import prisma from '~/lib/prisma'

/**
 * Checks if a user with the given email already exists.
 *
 * @param email The email to check
 * @param ignoreId The ID of the user to ignore
 * @returns Promise<boolean> Whether a user with the email already exists
 */
export default async function (email: string, ignoreId: number | null = null): Promise<boolean> {
    const clause: any = {
        where: { email },
    }

    if (ignoreId !== null) {
        clause.where.id = {
            not: ignoreId,
        }
    }

    const result = await prisma.user.count(clause)

    return result > 0
}
