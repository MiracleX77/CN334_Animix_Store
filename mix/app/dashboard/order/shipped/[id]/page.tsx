'use client'
import ProductComponent from "@/components/objects/product/Product"
import ProductFooter from "@/components/objects/product/ProductFooter"
import { useParams } from "next/navigation"
import OrderDetail from "@/components/objects/admin/admin/order/orderDetail"
import OrderProcessing from "@/components/objects/admin/admin/order/orderPorcessing"
import OrderShipped from "@/components/objects/admin/admin/order/orderShipped"

export default function ProductPage() {
    const params = useParams<{ id: string }>()
    return (
        <>
            <div className="container mx-auto py-10">
            <div className="">
                <div>
                <h1 className="text-2xl font-bold">Order Shipped</h1>
                <OrderShipped id={params.id} />
                </div>
                
            </div>
            
        </div>
        </>
    )
}