'use client'
import ProductComponent from "@/components/objects/product/Product"
import ProductFooter from "@/components/objects/product/ProductFooter"
import { useParams } from "next/navigation"

export default function ProductPage() {
    const params = useParams<{ product: string }>()
    return (
        <>
            <div className="mx-auto w-[80%] shadow-neon">

                <div className="p-4  rounded-xl">
                    <ProductComponent id={params.product} />
                    <ProductFooter id={params.product} />
                </div>

            </div>

        </>
    )
}