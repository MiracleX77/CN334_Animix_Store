'use client'
import OrderTable from "@/components/objects/admin/admin/order/orderTable";
import CardAdd from "@/components/objects/admin/admin/product/add/CardAdd";
import { useParams } from "next/navigation";


export default function DashboardOrderPage() {

  return (
    <div className="container mx-auto">
      <OrderTable/>
    </div>
  )
}