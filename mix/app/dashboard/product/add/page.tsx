'use client'
import CardAdd from "@/components/objects/admin/admin/product/add/CardAdd";
import { useParams } from "next/navigation";


export default function DashboardProductPage() {

  return (
    <div className="container mx-auto py-10">
      <h1 className="text-2xl font-bold"></h1>
      <CardAdd  />
    </div>
  )
}