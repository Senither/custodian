
export default function (event: any, status: number, reason: string, data: any = null): any {
    const response = { status, reason }

    setResponseStatus(event, status)

    return data ? { ...response, data } : response
}
