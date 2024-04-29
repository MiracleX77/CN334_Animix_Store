import { ProductCard } from "../ProductCard";
import CoverCarousel, { Cover } from "./ui/CoverCarousel";
import CoverHome from "./ui/CoverHome";

const coverArray: Cover[] = [
    {
        img: 'https://via.placeholder.com/600',
        title: 'Title 1',
        description: 'Description',
        button: 'Button',
        link: '/',
    },
    {
        img: 'https://via.placeholder.com/600',
        title: 'Title 2',
        description: 'Description',
        button: 'Button',
        link: '/',
    },
    {
        img: 'https://via.placeholder.com/600',
        title: 'Title 3',
        description: 'Description',
        button: 'Button',
        link: '/',
    },
    {
        img: 'https://via.placeholder.com/600',
        title: 'Title 4',
        description: 'Description',
        button: 'Button',
        link: '/',
    },
    {
        img: 'https://via.placeholder.com/600',
        title: 'Title 5',
        description: 'Description',
        button: 'Button',
        link: '/',
    },
]

export default function HomeCover() {
    return (
        <div className="w-full">
            <div className="p-4">
                <CoverCarousel cover={coverArray} />
            </div>
            
        </div>
    )

}