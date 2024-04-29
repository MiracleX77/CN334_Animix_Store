"use client"

import { useState } from "react"

type QuantityProps = {
    quantity: number
    setQuantity: (quantity: number) => void
}

const Quantity = ({ quantity, setQuantity }: QuantityProps) => {
    return (
        <div className="w-max rounded-md flex items-center border border-ring bg-card">
            <button
                onClick={() => setQuantity(quantity - 1)}
                className="w-10 h-14 flex justify-center items-center rounded-md hover:bg-accent hover:text-accent-foreground"
            >
                -
            </button>
            <input
                min={1}
                type="text"
                inputMode="numeric"
                value={quantity || 1}
                onChange={(e) => setQuantity(parseInt(e.target.value))}
                className="w-14 h-14 bg-card text-center flex justify-center items-center appearance-none"
            />
            <button
                onClick={() => setQuantity(quantity + 1)}
                className="w-10 h-14 flex justify-center items-center rounded-md hover:bg-accent hover:text-accent-foreground"
            >
                +
            </button>
        </div>
    )
}

export default Quantity
