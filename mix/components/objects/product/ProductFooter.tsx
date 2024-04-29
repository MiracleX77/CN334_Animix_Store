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
import {
    Tabs,
    TabsContent,
    TabsList,
    TabsTrigger,
} from "@/components/ui/tabs"
import { User, User2 } from "lucide-react"

export default function ProductFooter() {
    return (
        <>
            <div className="w-full p-4">
                <Tabs defaultValue="author" className="w-full">
                    <TabsList className="grid w-full grid-cols-2">
                        <TabsTrigger value="author">Author</TabsTrigger>
                        <TabsTrigger value="reviews">Reviews</TabsTrigger>
                    </TabsList>
                    <TabsContent value="author">
                        <Card>
                            <CardContent className="space-y-2">
                                <div className="mt-4 w-full h-full grid grid-cols-[150px_auto] items-center">
                                    <User2 size={100} />
                                    <div className="gap-2">
                                        <div className="flex space-x-2">
                                            <p>ผู้เเต่ง : </p>
                                            <p>อาซาโตะ อาซาโตะ </p>
                                        </div>
                                        <div className="flex space-x-2">
                                            <p>ภาพ : </p>
                                            <p>ชิราบิ </p>
                                        </div>
                                        <div className="flex space-x-2">
                                            <p>แปล  :  </p>
                                            <p>โชติกา ศรีภูริจรรยา </p>
                                        </div>
                                        <div className="flex space-x-2">
                                            <p>จำนวนหน้า  :  </p>
                                            <p>346 หน้า </p>
                                        </div>
                                    </div>
                                </div>
                            </CardContent>
                        </Card>
                    </TabsContent>
                    <TabsContent value="reviews">
                        <Card>
                            <CardHeader>
                                <CardTitle>Reviews</CardTitle>
                            </CardHeader>
                            <CardContent className="space-y-2">
                                <div className="mt-4 w-full h-full grid grid-cols-[150px_auto] items-center">
                                    <User2 size={100} />
                                    <div className="gap-2">
                                        <div className="flex space-x-2">
                                            <p>โดย : </p>
                                            <p>อาซาโตะ อาซาโตะ </p>
                                        </div>
                                        <div className="flex space-x-2">
                                            <p>รีวิว  :</p>
                                            <p>55555 </p>
                                        </div>
                                    </div>
                                </div>
                            </CardContent>
                        </Card>
                    </TabsContent>
                </Tabs>
            </div>
        </>
    )
}