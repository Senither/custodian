import { User } from '@prisma/client';
import { z } from 'zod';
import prisma from '~/lib/prisma';

export default function () {
    return {
        includeQuery: {
            select: {
                id: true,
                name: true,
                created_at: true,
            }
        },

        param: z.object({
            id: z.number({ coerce: true }).positive().int()
        }),

        createModelSchema: (user: User) => {
            return z.object({
                status: z.boolean().default(false).optional(),
                message: z.string().min(2).max(255),
                priority_id: z.number().int().positive(),
                category_id: z.number().int().positive(),
            })

                // Check if the priority ID provided is valid
                .refine(async (data) => {
                    return await prisma.priority.count({
                        where: {
                            id: data.priority_id,
                            user: user,
                        }
                    }) > 0
                }, {
                    message: 'The provided priority ID does not exist',
                    path: ['priority_id'],
                })

                // Check if the category ID provided is valid
                .refine(async (data) => {
                    return await prisma.category.count({
                        where: {
                            id: data.category_id,
                            user: user,
                        }
                    }) > 0
                }, {
                    message: 'The provided category ID does not exist',
                    path: ['category_id'],
                })
        },

        updateModelSchema: (user: User) => {
            return z.object({
                status: z.boolean().default(false).optional(),
                message: z.string().min(2).max(255).optional(),
                priority_id: z.number().int().positive().optional(),
                category_id: z.number().int().positive().optional(),
            })
                // Check if the priority ID provided is valid
                .refine(async (data) => {
                    if (!data.priority_id) {
                        return true
                    }

                    return await prisma.priority.count({
                        where: {
                            id: data.priority_id,
                            user: user,
                        }
                    }) > 0
                }, {
                    message: 'The provided priority ID does not exist',
                    path: ['priority_id'],
                })

                // Check if the category ID provided is valid
                .refine(async (data) => {
                    if (!data.category_id) {
                        return true
                    }

                    return await prisma.category.count({
                        where: {
                            id: data.category_id,
                            user: user,
                        }
                    }) > 0
                }, {
                    message: 'The provided category ID does not exist',
                    path: ['category_id'],
                })
        },
    }
}
