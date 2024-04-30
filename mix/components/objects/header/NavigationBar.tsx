'use client'

import React from "react"
import Link from "next/link"
import Image from "next/image"

import { cn } from "@/lib/utils"
import {
    NavigationMenu,
    NavigationMenuContent,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
    NavigationMenuTrigger,
    navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu"
import { FloatingLabelInput } from "@/components/ui/floatingInput"
import { Input } from "@/components/ui/input"
import { Heart, ShoppingCart, User2 } from "lucide-react"
import Cookies from "js-cookie"


export default function NavigationBar() {
    const role = Cookies.get("role")
    const token = Cookies.get("token")

    const [hasToken, setHasToken] = React.useState(token !== undefined)
    const [hasRole, setHasRole] = React.useState(role !== undefined)
    const [isAdmin,setIsAdmin] = React.useState(role === "Admin")

    const ListItem = React.forwardRef<
        React.ElementRef<"a">,
        React.ComponentPropsWithoutRef<"a">
    >(({ className, title, children, ...props }, ref) => {
        return (
            <li>
                <NavigationMenuLink asChild>
                    <a
                        ref={ref}
                        className={cn(
                            "",
                            className
                        )}
                        {...props}
                    >
                        <div className="text-sm font-medium leading-none">{title}</div>
                        <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
                            {children}
                        </p>
                    </a>
                </NavigationMenuLink>
            </li>
        )
    })
    ListItem.displayName = "ListItem"


    return (
        <>

            <nav className="w-full h-[70px] bg-Tretiary backdrop-blur-xl bg-opacity-20 dark:bg-opacity-60 text-md drop-shadow-lg">
                <div className="w-full h-full grid grid-cols-[100px_auto_200px] px-8">
                    <div className="flex items-center space-x-8">
                        <img src="/logo.png" alt="logo" className="h-8 w-8" />
                    </div>
                    <div className="w-full flex flex-col justify-center items-center">
                        <Input className="w-full p-1 ring-2 ring-opacity-50" placeholder="Search..."/>
                    </div>
                    <div className="w-full flex items-center justify-end space-x-8">
                        <div className="text-card-foreground">
                            {
                                hasToken && hasRole ? (
                                    <a href="/auth">
                                <User2 size={24} />
                            </a>
                                ) : ( 
                                    isAdmin ? (
                                        <a href="/dashboard">
                                <User2 size={24} />
                            </a>
                                    ) : (
                                        <a href="/profile">
                                <User2 size={24} />
                            </a>
                                )
                                )
                            }
                        </div>
                        <div className="text-card-foreground">
                            <a href="/cart">
                                <ShoppingCart size={24} />
                            </a>
                        </div>
                        <div className="text-card-foreground">
                            <a href="/favorite">
                                <Heart size={24} />
                            </a>
                        </div>
                    </div>
                </div>
            </nav>
        </>

    );
}