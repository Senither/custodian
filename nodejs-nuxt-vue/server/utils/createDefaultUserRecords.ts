import { Prisma, User } from '@prisma/client'
import prisma from '~/lib/prisma'

/**
 * Creates default records for a user.
 *
 * @param user The user to create default records for
 */
export default async function (user: User): Promise<void> {
    Promise.all([
        createDefaultCategoryRecords(user, [
            'House Stuff',
            'Work',
            'Learning',
            'Meeting',
        ]),
        createDefaultPriorityRecords(user, [
            'Low',
            'Medium',
            'High',
            'Highest',
        ]),
    ]).catch(console.error)
}

async function createDefaultCategoryRecords(user: User, categories: string[]): Promise<Prisma.BatchPayload> {
    return prisma.category.createMany({
        data: categories.map(name => ({
            user_id: user.id,
            name,
        })),
    })
}

async function createDefaultPriorityRecords(user: User, priorities: string[]): Promise<Prisma.BatchPayload> {
    return prisma.priority.createMany({
        data: priorities.map(name => ({
            user_id: user.id,
            name,
        })),
    })
}
