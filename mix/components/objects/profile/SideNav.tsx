'use client'
import Link from "next/link"
import {
    Bell,
    Home,
    LineChart,
    LogOut,
    Package,
    Package2,
    ShoppingCart,
    Store,
    Users,
} from "lucide-react"

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { getUser } from "@/apis/services/userService"
import { useEffect, useState } from "react"

export default function SideNav() {
    const [user, setUser] = useState<any>(null)
    const [openSheet, setOpenSheet] = useState(false)
    const [isLoad, setIsLoad] = useState(true)


    useEffect(() => {
        loadProducts().then(() => {

        }).catch((error) => {
            console.error("An error occurred while fetching data:", error);
        });
    }, []);

    const loadProducts = async () => {
        const token = AuthProvider.getAccessToken()
        await getUser(token||'').then((res) => {
            setUser(res.data.data);
            console.log(res.data.data);
            setIsLoad(false)
        }).catch((error) => {
            console.error("An error occurred while fetching data:", error);
        });
    };
    return (
        <>
        {
            !isLoad && (
                <div className="grid min-h-screen w-full grid-rows-[auto_90px]">
                <div className="hidden border-r bg-muted/40 md:block">
                    <div className="flex h-full max-h-screen flex-col gap-2">
                        <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
                            <Link href="/" className="flex items-center gap-2 font-semibold">
                                <Store className="h-6 w-6" />
                                <span className="">{user.first_name} {user.last_name}</span>
                            </Link>
                        </div>
                        <div className="flex-1">
                            <nav className="grid items-start px-2 text-sm font-medium lg:px-4">
                                <Link
                                    href="/profile/edit"
                                    className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                                >
                                    <Home className="h-4 w-4" />
                                    Account
                                </Link>
                                <Link
                                    href="/profile/orders"
                                    className="flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all hover:text-primary"
                                >
                                    <Package className="h-4 w-4" />
                                    Order
                                </Link>
                                <Link
                                    href="/profile/address"
                                    className="flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all hover:text-primary"
                                >
                                    <Package className="h-4 w-4" />
                                    Address{" "}
                                </Link>
                                <Link
                                    href="/profile/review"
                                    className="flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all hover:text-primary"
                                >
                                    <Package className="h-4 w-4" />
                                    Review{" "}
                                </Link>
                            </nav>
                        </div>
                    </div>
                </div>
                <div>
                    <Link
                        href="#"
                        className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
                    >
                        <LogOut className="h-4 w-4" />
                        Logout
                    </Link>
                </div>
            </div>
            )
        }
        </>
    )
}
