import { User } from '@prisma/client'
import { z } from 'zod'
import prisma from '~/lib/prisma'

export default defineEventHandler(async (event) => {
    const { user, err } = await loadAuthenticatedUserFromSession(event)
    if (err) {
        return err
    }

    const validatedBody = await readValidatedBody(event, createValidationSchema(user).safeParseAsync)
    if (!validatedBody.success) {
        return createErrorResponse(event, 400, 'Validation failed', validatedBody.error)
    }

    await prisma.user.update({
        where: { id: user.id },
        data: validatedBody.data,
    })

    return {
        status: 200,
        message: 'User details updated successfully',
        data: validatedBody.data
    }
})

function createValidationSchema(user: User) {
    return z.object({
        name: z.string().min(2).optional(),
        email: z.string().email().optional(),
        current_password: z.string().optional(),
        password: z.string().min(8).optional(),
        password_confirmation: z.string().optional(),
    })
        // Add check to ensure the email is unique if it's being updated, with the exception of the current user's email
        .refine(async data => {
            if (!data.email) {
                return true
            }
            return !await checkUserWithEmailAlreadyExists(data.email, user.id)
        }, {
            message: 'Email has already been taken',
            path: ['email'],
        })

        // Add check to ensure the password and password_confirmation fields match when updating the password
        .refine(async data => {
            if (!data.password) {
                return true
            }

            if (!data.password_confirmation) {
                return false
            }

            return data.password === data.password_confirmation
        }, {
            message: 'Password and password confirmation do not match',
            path: ['password'],
        })

        // Add check to ensure the current_password is provided if the password is being updated
        .refine(data => {
            if (!data.password) {
                return true
            }

            return !!data.current_password
        }, {
            message: 'Current password is required to update password',
            path: ['current_password'],
        })

        // Add check to ensure that the current_password is valid and matches the user's current password
        .refine(async data => {
            if (!data.current_password) {
                return true
            }

            return await verifyPassword(user.password, data.current_password)
        }, {
            message: 'Current password is incorrect',
            path: ['current_password'],
        })

        // Transform the data to remove verification fields
        .transform(async data => {
            if (data.current_password && data.password && data.password_confirmation) {
                data.password = await hashPassword(data.password)
            }

            delete data.password_confirmation
            delete data.current_password

            return data
        })
}
