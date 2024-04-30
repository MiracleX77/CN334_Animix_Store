import SideNav from "@/components/objects/profile/SideNav";

export default function ProfileLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="w-full h-full">
            <div className="h-full w-full grid md:grid-cols-[220px_auto] lg:grid-cols-[280px_auto]">
                <SideNav />
                <div className={`w-full h-full flex flex-col bg-Bg text-Black overflow-y-auto`}>
                    <div className='w-full h-full'>{children}</div>
                </div>
            </div>
        </div>
    )
}