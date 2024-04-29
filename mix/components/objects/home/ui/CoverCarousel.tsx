'use client'
import { type CarouselApi } from "@/components/ui/carousel"
import {
    Carousel,
    CarouselContent,
    CarouselItem,
} from "@/components/ui/carousel"
import Autoplay from "embla-carousel-autoplay"
import { useEffect, useState } from "react"
import CoverHome from "./CoverHome"
import Middle from "@/components/layouts/Middle"

export interface Cover {
    img: any;
    title: string;
    description: string;
    button: string;
    link: string;
}

type Props = {
    cover: Cover[];
}

export default function CoverCarousel({ cover }: Props) {

    const [carouselApi, setCarouselApi] = useState<CarouselApi>()

    //get mouse enter and leave event
    const [mouse, setMouse] = useState(false)
    const [current, setCurrent] = useState(0)
    const [total, setTotal] = useState(0)
    const handleMouseEnter = () => { setMouse(true) }
    const handleMouseLeave = () => { setMouse(false) }

    useEffect(() => {
        if (carouselApi) {
            setTotal(carouselApi.scrollSnapList().length)
            carouselApi.on("select", () => {
                setCurrent(carouselApi.selectedScrollSnap())
            })
        }
    }, [carouselApi])

    return (
        <>
            <Carousel className="w-full shadow-neon p-1 rounded-2xl"
                opts={
                    {
                        align: "start",
                        loop: true,
                    }
                }
                plugins={[Autoplay({
                    delay: 5000,
                    stop: mouse,
                    reset: !mouse,
                })]}
                setApi={setCarouselApi}
            >
                <CarouselContent className="-ml-1" >
                    {Array.from({ length: cover.length }).map((_, index) => (
                        <CarouselItem key={index} className="pl-1"
                            onMouseEnter={handleMouseEnter}
                            onMouseLeave={handleMouseLeave}
                        >
                            <CoverHome
                                cover={cover[index]}
                            />
                        </CarouselItem>
                    ))}
                </CarouselContent>
            </Carousel>
            {/* <Middle X Y className="mt-2 w-full space-x-2">
                {Array.from({ length: total }).map((_, index) => (
                    <button key={index} onClick={() => carouselApi?.scrollTo(index)}
                        className={`w-2 h-2 rounded-full ${current === index ? 'bg-primary w-6' : 'bg-Gray-light'} animate`}
                    />
                ))}
            </Middle> */}
        </>
    )
}