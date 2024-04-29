'use client'
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function ButtonAdd() {
    const router = useRouter()
    const addProduct = () => {
        router.push('/dashboard/product/add')
    }
    return (
        <Button onClick={addProduct} className="bg-primary ring-2 px-8 rounded-l text-white">
            Add
        </Button>
    )
}