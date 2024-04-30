'use client'
import AddressAccount from "@/components/objects/profile/address/addressAccount"
import AddressAdd from "@/components/objects/profile/address/addressAdd"


export default function AddressAccountPage() {


return (
        <div className="container mx-auto py-10">
            <div className="flex justify-between items-center mb-4">
                <div>
                <h1 className="text-2xl font-bold">Address</h1>
                <h2 className="text-lg">Manage your address</h2>
                </div>
                <AddressAdd />
            </div>
            <AddressAccount />
        </div>
)
}