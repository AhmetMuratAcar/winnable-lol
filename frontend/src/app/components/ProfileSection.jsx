"use client"
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { regionTagToServerCode } from "../utils/idValidation";

export default function ProfileSection({ region, slug }) {
    const router = useRouter()
    const [data, setData] = useState(null)
    const [loading, setLoading] = useState(true)
    // const [errorMessage, setErrorMessage] = useState('')
    const regionServerCode = regionTagToServerCode(region)

    useEffect(() => {
        if (typeof slug !== 'string' || !slug.includes('-')) {
            setLoading(false)
            router.push('/summoner-not-found')
            return
        }

        const [gameName, tagLine] = slug.split('-') || []
        if (!gameName || !tagLine) {
            setLoading(false)
            router.push('/summoner-not-found')
            return
        }

        async function fetchData() {
            try {
                const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/lol/profile`, {
                    method: `POST`,
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        region: regionServerCode,
                        gameName: gameName,
                        tagLine: tagLine,
                    }),
			    })

                if (!res.ok) {
                    if (res.status === 404) {
                        router.push('/summoner-not-found')
                    } else {
                        router.push(`/server-error?status=${res.status}`)
                    }
                    setLoading(false)
                    return
			    }

                const json = await res.json()
                setData(json)
            } catch (err) {
                console.log(`ERROR: ${err}`)
                const isConnectionRefused = err instanceof TypeError && err.message.includes('Failed to fetch')
                
                if (isConnectionRefused) {
                    console.log('our servers are down')
                    // route to our servers are down page
                } else {
                    console.log('unexpected error')
                    // route to unexpected error page
                }
            } finally {
                setLoading(false)
            }
        }
        fetchData()
    }, [region, slug])

    if (loading) {
        // render a loading icon on page
        return <p>Loading…</p>
    }

    return (
        <section 
            id="ProfileSection"
            className="flex flex-col flex-grow items-center justify-start px-4 py-6"
        >
            <div>
                <p>profile section rendered</p>
                <div className="w-full max-w-lg space-y-4">
                {data && data.length > 0 ? (
                    data.map(({ championId, championLevel, championPoints }) => (
                    <div
                        key={championId}
                        className="p-4 border rounded-md bg-[var(--contrast)] text-white"
                    >
                        <p>Champion ID: {championId}</p>
                        <p>Level: {championLevel}</p>
                        <p>Points: {championPoints}</p>
                    </div>
                    ))
                ) : (
                    <p>No mastery data to show.</p>
                )}
                </div>
            </div>
        </section>
    );
}