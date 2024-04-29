'use client'
import CardEdit from "@/components/objects/admin/admin/product/edit/CardEdit";
import { useParams } from "next/navigation";


export default function DashboardProductPage() {
  const params = useParams<{ product: string }>()

  return (
    <div className="container mx-auto py-10">
      <h1 className="text-2xl font-bold">{params.product}</h1>
      <CardEdit id={params.product} />
    </div>
  )
}