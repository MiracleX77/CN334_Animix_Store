import { Button } from "@/components/ui/button";

export default function Category() {

    interface ButtonProps {
        title: string;
        href: string;
    }

    const button = [
        {
            title: "Best Seller",
            href: "/"
        },
        {
            title: "New Realease",
            href: "/category/newrealease"
        },
        {
            title: "Flash Sale",
            href: "/category/flashsale"
        },
        {
            title: "Manga",
            href: "/category/manga"
        },
        {
            title: "Light Novel",
            href: "/category/light-novel"
        },
        {
            title: "Yaoi",
            href: "/category/yaoi"
        },
        {
            title: "Yuri",
            href: "/category/yuri"
        },
        {
            title: "Artbook",
            href: "/category/artbook"
        }
        
    ]
    return (
        <>
            <div className="h-[70px] w-full bg-background flex justify-between items-center">
                {
                    button.map((item: ButtonProps, index: number) => (
                        <Button key={index} className="text-card-foreground bg-transparent ring" ><a href={item.href}>{item.title}</a></Button>
                    ))
                }
            </div>
        </>
    )
}