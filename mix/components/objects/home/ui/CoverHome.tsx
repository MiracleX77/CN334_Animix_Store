import Image from 'next/image';
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import Link from 'next/link';
import { Cover } from './CoverCarousel';

type Props = {
    cover: Cover;
}

export default function CoverHome({ cover }: Props) {
    return (
        <div className='relative w-full'>
            <div className='aspect-[2/1] overflow-hidden relative'>
                <Image
                    src={cover.img} alt={'cover'} layout='cover' width={1920} height={1080}
                    className='w-full aspect-[2/1] object-cover rounded-2xl'
                />
            </div>
            <Card className={cn("z-10 mx-auto mt-4 absolute right-[5%] bottom-10 rounded-full border-2 bg-black bg-opacity-50")}>
                <div className='p-2 bg-opacity-40'>
                    <Link href={cover.link}>
                        <p className=" text-white">{`${cover.title} >>`}</p>
                    </Link>
                </div>
            </Card>


        </div>
    );
}
