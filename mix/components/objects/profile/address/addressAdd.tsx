'use client'
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function AddressAdd() {
    const router = useRouter()
    const addAddress = () => {
        router.push('/profile/address/add')
    }
    return (
        <Button onClick={addAddress} className="bg-primary ring-2 px-8 rounded-l text-white">
            Add
        </Button>
    )
}