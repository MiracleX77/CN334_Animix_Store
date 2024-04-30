'use client'
import React, { useState, useEffect } from 'react';
import { DataTable } from './data-table';
import { getProductById,deleteProductById } from "@/apis/services/productService"
import { getAddressAll } from '@/apis/services/addressService';
import { AuthProvider } from "@/utils/clientAuthProvider"
import { Product ,columns} from './columns';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { useRouter } from "next/navigation"
import { set } from 'date-fns';
import Cookies from 'js-cookie';
import { Button } from '@/components/ui/button';

interface ProductIdCount {
    [key: string]: number;
}

function CartTable() {

    const router = useRouter();
    const token = AuthProvider.getAccessToken()

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertOnCancel, setAlertOnCancel] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    const [alertCancelBottom, setAlertCancelBottom] = useState(false);

    const [products, setProducts] = useState<Product[]>([]);
    const [total, setTotal] = useState(0);

    const[isLoad,setIsLoad] = useState(true);

    useEffect(() => {
        loadProducts().then(() => {
            setIsLoad(false);
        }).catch((error) => {
            console.error("An error occurred while fetching data:", error);
            setIsLoad(false);
        });
    }, []);

    const loadProducts = async () => {
        const cart = Cookies.get("cart");
        const token = AuthProvider.getAccessToken();

        const cart_a = cart?.slice(1, -1);
        const product_id = cart_a?.split(",");
        
        console.log("Product ID:", product_id);
        if (product_id !== null && product_id !== undefined) {
            try {
                const productIdCount: ProductIdCount = product_id.reduce((countObj: ProductIdCount, id: string) => {
                    countObj[id] = (countObj[id] || 0) + 1;
                    return countObj;
                }, {});
                console.log("Product ID count:", productIdCount);
                // Use Promise.all to fetch all products concurrently
                const productPromises = Object.keys(productIdCount).map((id) => getProductById(id, token || ''));
                const products = await Promise.all(productPromises);
                const newProducts = products.map((productRes) => {
                    const product = productRes.data.data;
                    product.count = productIdCount[product.id];
                    product.price = product.price * product.count;
                    product.onRemove = () => handleDelete(product.id);
                    return product;
                });
                const newTotal = newProducts.reduce((acc, product) => acc + product.price, 0);
                setTotal(newTotal);
                setProducts(newProducts);
  // Save data to a cookie
            } catch (error) {
                console.error("Failed to load products:", error);
            }
        }
    };
    

    const handleDelete = async (id:string) => {
        const cart = Cookies.get("cart");
        const cartIds: string[] = cart ? JSON.parse(cart) : [];
        const newCartIds = cartIds.filter((cartId) => cartId !== id);
        Cookies.set("cart", JSON.stringify(newCartIds), { expires: 7 });
        setOpenAlert(false);
        loadProducts()
    };

    const handleCheckout = async () => {

        const res = await getAddressAll(token||'');
        console.log(res);
        if (res.data.data.length === 0) {
            setAlertTitle('No Address');
            setAlertContent('Please add your address before checkout');
            setAlertStatus('error');
            setAlertConfirmText('OK');
            setAlertOnConfirm(() => () => router.push('/profile/address'));
            setOpenAlert(true);
            return;
        }
        const newProducts = products.map((product) => {
            return {
                id: product.id,
                count: product.count,
                price: product.price
            };
        });
        const dataString = JSON.stringify(newProducts);
        Cookies.set('checkoutData', dataString);
        router.push('/cart/checkout');
    }
   

  return (
    <>{
        !isLoad && ( 
            <div>
                <AlertDialog 
            open={openAlert} 
            setOpen={setOpenAlert} 
            title={alertTitle} 
            content={alertContent} 
            status={alertStatus} 
            onConfirm={alertOnConfirm} 
            onCanceled={alertOnCancel}
            confirmText={alertConfirmText} 
            cancelBottom={alertCancelBottom}
        />
        <DataTable columns={columns} data={products} />
        <div className="flex justify-between mt-4 mb-6">
            <div className="flex flex-col">
                <p className="text-lg">Total</p>
                <p className="text-lg">Shipping</p>
                <p className="text-lg">Subtotal</p>
            </div>
                <div className="flex flex-col items-end">
                    <p className="text-lg">฿ {total}</p>
                    <p className="text-lg">฿ 0</p>
                    <p className="text-lg">฿ {total}</p>
                </div>
        </div>
        <div className="flex justify-center"> {/* Centering button */}
        <Button onClick={handleCheckout} className="bg-primary ring-2 px-10 py-2 rounded-xl text-white">
            Checkout
        </Button>
        </div>
            </div>
            
        )

    }
    
        
    </>
    
  );
}

export default CartTable;