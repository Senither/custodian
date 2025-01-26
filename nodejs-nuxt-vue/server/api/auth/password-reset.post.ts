import { z } from 'zod'
import prisma from '~/lib/prisma'

const schema = z.object({
    email: z.string().email(),
}).refine(async data => await checkUserWithEmailAlreadyExists(data.email), {
    message: 'We can\'t find a user with that email address.',
    path: ['email'],
})

export default defineEventHandler(async (event) => {
    const validatedBody = await readValidatedBody(event, schema.safeParseAsync)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Missing required fields', validatedBody.error)
    }

    // Delete the existing password reset tokens for this email address
    await prisma.passwordResetToken.deleteMany({
        where: {
            email: validatedBody.data.email,
        },
    })

    // Create a new password reset token
    await prisma.passwordResetToken.create({
        data: {
            email: validatedBody.data.email,
            token: generateToken(),
        },
    })

    // Here is where we'd send the email to the user with the password
    // reset link, this could be done with something like:
    // https://nuxt.com/modules/nuxt-mail
    // However, for this example we will just return a success message.

    return {
        statusCode: 200,
        message: 'We have emailed your password reset link.',
    }
})

/**
 * Generates a random alpha-numeric token with the given length.
 *
 * @param length The length of the token to generate
 * @returns The generated token
 */
function generateToken(length: number = 64): string {
    const charset = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'

    return Array.from({ length })
        .map(() => charset.charAt(Math.random() * charset.length))
        .join('')
}
