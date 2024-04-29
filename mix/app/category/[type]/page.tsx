'use client'
import ListProductCard from "@/components/objects/ListProductCard"
import { useParams } from "next/navigation"

export default function ProductPage() {
    const params = useParams<{ type: string }>()
    //  make manga -> Manga
    params.type = params.type.charAt(0).toUpperCase() + params.type.slice(1)
    if (params.type === 'Manga') {
        params.type = 'Manga'
    } else if (params.type === 'Light-novel') {
        params.type = 'Light Novel'
    }


    return (
        <>
            <div className="mx-auto w-[80%] px-2 md:px-4 lg:px-6">
            <h1 className="text-2xl font-bold">{params.type}</h1>
            <div className="mt-4 p-2 md:p-4 lg:p-6 ring-4 rounded-xl">
                <ListProductCard type={params.type} />
            </div>
            </div>

        </>
    )
}