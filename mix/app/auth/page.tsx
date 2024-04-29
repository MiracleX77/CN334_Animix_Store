import Middle from "@/components/layouts/Middle";
import Auth from "@/components/objects/auth/autx";

export default function AuthPage() {
    return (
        <div className="w-screen h-screen">
            <Middle X Y className="w-full h-full">
                <Auth />
            </Middle>
        </div >
    )
}