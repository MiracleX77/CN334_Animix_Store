'use client'
import React, { useState } from "react"
import Image from "next/image"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import Middle from "../layouts/Middle"
import Span from "../layouts/Span"

export interface ProductCard {
    img_url: string;
    name: string
    description?: string
    price: number
    onClick?: () => void
    new?: boolean
    hot?: boolean
    sale?: number
}

type props = {
    product: ProductCard
}

export function ProductCard({ product }: props) {
    
    const [hover, setHover] = useState(false)

    return (
        <>
            <Card className="w-[120px] sm:w-[150px] md:w-[200px] 2xl:w-[300px] p-2 pb-0 cursor-pointer shadow-lg rounded-xl ring-4"
                onClick={() => setHover(true)}
            >
                <div className="relative overflow-hidden rounded-xl">
                    <img src={product.img_url} alt="Product" width={360} height={360} className="rounded-lg h-[200px] sm:h-[240px] md:h-[300px] 2xl:h-[360px] hover:scale-110 animate" />
                </div>
                <div className="flex justify-between mt-2 p-2 gap-4">
                    <div className="w-full">
                        <CardTitle>{product.name}</CardTitle>
                        <CardDescription>
                            <p className="text-lg md:text-xl text-green-500 text-right">฿ {product.price}</p>
                        </CardDescription>
                    </div>
                    <Button className="bg-secondary ring-2 px-4 rounded-xl text-white">
                        Add
                    </Button>
                </div>
            </Card>
        </>

    )
}
