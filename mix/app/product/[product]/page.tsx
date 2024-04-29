'use client'
import ProductComponent from "@/components/objects/product/Product"
import ProductFooter from "@/components/objects/product/ProductFooter"
import { useParams } from "next/navigation"

export default function ProductPage() {
    const params = useParams<{ product: string }>()
    return (
        <>
            {params.product}
            <div className="w-full p-4">

                <div className="p-4 ring-4 rounded-xl">
                    <ProductComponent id={params.product} />
                    <ProductFooter />
                </div>

            </div>

        </>
    )
}