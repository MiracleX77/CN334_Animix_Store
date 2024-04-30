'use client'
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
import { useEffect, useState } from "react"
import { getProductById} from "@/apis/services/productService"
import { getReviewsByProductId,postReview } from "@/apis/services/reviewService"
import { ReviewCreate } from "@/models/dto/review"
import { get } from "http"
import { useRouter } from "next/navigation"
import { AuthProvider } from "@/utils/clientAuthProvider"
import AlertDialog from "@/components/interactive/layout/alertdialog"


export default function ProductFooter({ id }: { id: string }) {

    const router = useRouter();
    const token = AuthProvider.getAccessToken();
    const [product, setProduct] = useState<any>(null);
    const [loading, setLoading] = useState(true);
    const [reviews, setReviews] = useState<any[]>([]);

    const [hasReview, setHasReview] = useState(false);
    const [hasToken, setHasToken] = useState(false);

    const [rating, setRating] = useState("");
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');


    const handleRating = (e: any) => {
        setRating(e.target.value);
    }
    const handleTitle = (e: any) => {
        setTitle(e.target.value);
    }
    const handleContent = (e: any) => {
        setContent(e.target.value);
    }

    const handleSubmit = async (event:any) => {
        event.preventDefault();
        console.log("submit");
        const data: ReviewCreate = {
            product_id: parseInt(id),
            title: title,
            content: content,
            rating:  parseInt(rating),
        }
        console.log(data.rating);

        await postReview(data,token || "").then((res) => {
            console.log(res);
            if (res) {
                if (res.status === 200) {
                    console.log("success");
                    setAlertTitle("Review");
                    setAlertContent("Review added successfully");
                    setAlertConfirmText("OK");
                    setAlertStatus("success");
                    setAlertOnConfirm(() => () => setOpenAlert(false));
                    setOpenAlert(true);
                } else {
                    console.log("fail");
                    setAlertTitle("Review");
                    setAlertContent("Review added failed");
                    setAlertConfirmText("OK");
                    setAlertStatus("error");
                    setAlertOnConfirm(() => () => setOpenAlert(false));
                    setOpenAlert(true);
                }
            }
        });

    }


    

    useEffect(() => {
        getProductById(id, "").then((res) => {
            res.data.data.created_at = new Date(res.data.data.created_at).toLocaleDateString();
            setProduct(res.data.data);
            setLoading(false);
        });
        getReviewsByProductId(id, "").then((res) => {
            const data = res.data.data;
            data.forEach((review: any) => {
                review.created_at = new Date(review.created_at).toLocaleDateString();
            });
            setReviews(data);
        });
        const token = AuthProvider.getAccessToken();
        if (token) {
            setHasToken(true);
        }
    }, []);

    return (
        <>
            <AlertDialog open={openAlert} setOpen={setOpenAlert} title={alertTitle} content={alertContent} status={alertStatus} onConfirm={alertOnConfirm} confirmText={alertConfirmText} cancelBottom={false}/>

            {!loading && (<div className="w-full ">
                <Tabs defaultValue="details" className="w-full">
                    <TabsList className="grid w-full grid-cols-3">
                        <TabsTrigger value="details">Details</TabsTrigger>
                        <TabsTrigger value="reviews">Reviews</TabsTrigger>
                        <TabsTrigger value="write">Write-Reviews</TabsTrigger>
                    </TabsList>
                    <TabsContent value="details">
                        <Card>
                            <CardContent className="space-y-2">
                                <div className="mt-4 w-full h-full grid grid-cols-[150px_auto] items-center">
                                    <User2 size={100} />
                                    <div className="gap-2">
                                        <div className="flex space-x-2">
                                            <p>Author : </p>
                                            <p>{product.author.name}</p>
                                        </div>
                                        <div className="flex space-x-2">
                                            <p>Description : </p>
                                            <p>{product.author.description} </p>
                                        </div>
                                        <div className="flex space-x-2">
                                            <p>Publisher :  </p>
                                            <p>{product.publisher.name} </p>
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
                                {reviews.map((review) => (
                                    <div key={review.id} className="mt-4 w-full grid grid-cols-[auto_1fr] items-start gap-4">
                                    <div className="place-self-start">
                                        <User2 size={60} />
                                    </div>
                                    <div>
                                        <div className="flex justify-between items-center">
                                            <p className="text-gray-200 font-semi">{review.first_name} {review.last_name}</p>
                                            <span className=" text-yellow-400  mr-1">{Array(review.rating).fill('★').join('')}</span>
                                        </div>
                                        <div className="flex justify-between items-center">

                                            <p className={`font-semibold ${review.polarity == "positive" ? 'text-green-500' : 'text-red-500'}`}>
                                                {review.title}</p>
                                            <p className="text-gray-400">{review.created_at}</p>
                                        </div>
                                        <p className="text-gray-400">{review.content}</p>
                                    </div>
                                    </div>
                                ))}
                                </CardContent>
                        </Card>
                    </TabsContent>
                    <TabsContent value="write">
                        <Card>
                            <CardHeader>
                                <CardTitle>Reviews</CardTitle>
                            </CardHeader>
                            {
                                hasToken ? (
                                    <CardContent className="space-y-2">
                                        <form onSubmit={handleSubmit}>
                                        <div className="flex flex-col col-span-2 row-span-2">
                                        <div className="flex items-center space-x-4">
                                            <div className="flex flex-col ">
                                            <Label>Title</Label>
                                            <Input type="text" required id="title" value={title} onChange={handleTitle} />
                                            </div>
                                            <div className="flex flex-col">
                                            <Label>Rating</Label>
                                            <Input type="number" onChange={handleRating} required min={1} max={5} />
                                            </div>
                                        </div>
                                        <div className="flex flex-col">
                                            <Label>Content</Label>
                                            <textarea onChange={handleContent} required className="form-textarea mt-1 block w-full" rows={5}></textarea>
                                        </div>
                                            <Button type="submit" className="bg-primary ring-2 px-4 py-2 rounded-xl text-white w-full sm:w-auto">
                                                Submit
                                            </Button>
                                        </div>
                                        </form>
                                    </CardContent>
                                ) : (
                                    <CardContent className="flex flex-col items-center justify-center space-y-2 mx-auto h-full">
                                        <CardDescription className="text-center">
                                            You must be logged in to write a review.
                                        </CardDescription>
                                        <CardFooter>
                                            <Button onClick={()=>router.push("/auth")} className="bg-primary ring-2 px-10 rounded-xl text-white ">
                                                Login
                                            </Button>
                                        </CardFooter>
                                    </CardContent>
                                )
                            }
                        </Card>
                    </TabsContent>
                </Tabs>
            </div>
            )}
        </>
    )
}