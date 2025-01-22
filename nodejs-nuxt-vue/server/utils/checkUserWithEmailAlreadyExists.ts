import prisma from '~/lib/prisma'

/**
 * Checks if a user with the given email already exists.
 *
 * @param email The email to check
 * @returns Promise<boolean> Whether a user with the email already exists
 */
export default async function (email: string): Promise<boolean> {
    const user = await prisma.user.findFirst({
        where: { email },
    })

    return user !== null
}
