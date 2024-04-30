import CheckOutTable from "@/components/objects/cart/checkout/cartTable"
import CheckOut from "@/components/objects/cart/checkout/checkout"

export default function CheckOutPage() {
    return (
        <>
            <div className="mx-auto w-[80%] shadow-neon mt-5">
                <div className="p-4  rounded-xl">
                <CheckOut></CheckOut>
                </div>
            </div>

        </>
    )
}