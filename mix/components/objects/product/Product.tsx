'use client'
import { Button } from '@/components/ui/button';
import Quantity from '@/components/ui/quantity';
import Image from 'next/image';
import { use, useEffect, useState } from 'react';
import { getProductById} from '@/apis/services/productService';
import { getReviewsByProductId } from '@/apis/services/reviewService';
import { get } from 'http';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import Cookies from 'js-cookie';


export default function ProductComponent({ id }: { id: string }) {
    
    const [product, setProduct] = useState<any>(null);
    const [reviews, setReviews] = useState<any[]>([]);
    const [loading, setLoading] = useState(true);
    const [quantity, setQuantity] = useState(1);

    const [countReview, setCountReview] = useState(0);
    const [allRating, setAllRating] = useState(0);

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');

    useEffect(() => {
        Promise.all([
            getProductById(id, ""),
            getReviewsByProductId(id, "")
        ]).then(([productRes, reviewsRes]) => {
        const productData = productRes.data.data;
        productData.created_at = new Date(productData.created_at).toLocaleDateString();
        setProduct(productData);

        const reviewsData = reviewsRes.data.data;
        setCountReview(reviewsData.length);

        const allRating = reviewsData.reduce((acc: any, review: { rating: any; }) => acc + review.rating, 0);
        setAllRating(reviewsData.length > 0 ? allRating / reviewsData.length : 0);
    }).catch((error) => {
        console.error("An error occurred while fetching data:", error);
    }).finally(() => {
        setLoading(false);
    });
    }, []);

    const handleAddToCart = () => {
        console.log("Add to cart");
        let cart = Cookies.get("cart");
        let cartIds: string[] = cart ? JSON.parse(cart) : [];
        for (let i = 0; i < quantity; i++){
            cartIds.push(product.id);
        }
        Cookies.set("cart", JSON.stringify(cartIds), { expires: 7 });  
        console.log(`Product ID ${product.id} added to cart`);
        setAlertTitle("Add to Cart");
        setAlertContent("Product added to cart successfully");
        setAlertConfirmText("OK");
        setAlertStatus("success");
        setAlertOnConfirm(() => () => setOpenAlert(false));
        setOpenAlert(true);
    }


    console.log(product);

    return (
        <>
            <AlertDialog open={openAlert} setOpen={setOpenAlert} title={alertTitle} content={alertContent} status={alertStatus} onConfirm={alertOnConfirm} confirmText={alertConfirmText} cancelBottom={false}/>

            {!loading && (
            <div className="w-full h-full grid grid-cols-1 md:grid-cols-2 gap-4"> 
                <div>
                    <div className="m-4 before:shadow-neon"> {/* Added margin */}
                        <img src={product.img_url} alt="product" className='w-full h-[250px] sm:h-[450px] md:h-[550px] 2xl:h-[700px] rounded-2xl' />
                    </div>
                </div>
                <div className="m-3"> {/* Added margin */}
                    <h2 className='text-4xl font-bold mb-4'> {/* Added bottom margin */}{product.name}</h2>
                    <p className='text-sm text-secondary-foreground mb-2'>วันวางจำหน่าย : {product.created_at}</p> {/* Added bottom margin */}
                    <div className="flex items-center text-yellow-400 text-lg mb-7">
                        <span className=" text-yellow-400  mr-1">{Array(Math.floor(allRating)).fill('★').join('')}</span>
                        <span className="text-white text-sm">{allRating}</span>
                        <span className="text-white text-sm ml-1 mr-10"> / 5.0</span>
                        <span className="text-white text-sm">{countReview} review</span>
                    </div>
                    <p className='text-2xl mb-4 text-green-400'>฿ {product.price} </p> {/* Added bottom margin */}
                    <div className='flex items-center space-x-4 mb-6'> {/* Added bottom margin */}
                        <p>Count :</p>
                        <Quantity quantity={quantity} setQuantity={setQuantity} />
                    </div>
                    <Button onClick={handleAddToCart} className="relative text-white py-2 px-4 rounded-xl bg-green-700 before:absolute before:-inset-1 before:-z-10 before:block before:rounded-xl before:bg-blue-500 before:shadow-neon">
                        Add
                    </Button>

                    <div className="mt-8"> {/* Added top margin */}
                        <h2 className='text-xl font-bold mb-2'>Description</h2> {/* Added bottom margin */}
                        <p>
                            {product.description}
                        </p>
                    </div>
                </div>
            </div>
        )}
        </>
    )
}

