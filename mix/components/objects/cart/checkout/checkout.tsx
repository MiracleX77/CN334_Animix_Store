"use client"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"

import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { getAddressAll} from "@/apis/services/addressService"
import { ChangeEvent, use, useEffect, useState } from "react"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { ProductModel } from "@/models/dto/product"
import { set } from "date-fns"
import { Value } from "@radix-ui/react-select"
import Image from "next/image"
import { Button } from '@/components/ui/button';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { ProductCreate } from "@/models/dto/product"
import { useRouter } from "next/navigation"
import { AddressCreate,Address } from "@/models/dto/address"
import CheckOutTable from "./cartTable"
import Cookies from 'js-cookie';
import { OrderCreate } from "@/models/dto/order"
import { postOrder } from "@/apis/services/orderServices"

export default function CheckOut() {
    const token = AuthProvider.getAccessToken()
    const router = useRouter();

    const [showQRModal, setShowQRModal] = useState(false);
    const [image, setImage] = useState<File | null>(null);

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertOnCancel, setAlertOnCancel] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    const [alertCancelBottom, setAlertCancelBottom] = useState(false);

    const [isLoad,setIsLoad] = useState(true);

    type ProductModel = {
        id:number,
        price:number,
        count:number,
    }
    const [products, setProducts] = useState<ProductModel[]>([]);

    const [address, setAddress] = useState<Address[]>([]);
    const [selectedAddress, setSelectedAddress] = useState<Address>({
        id: 0,
        user_id: 0,
        address_line: '',
        phone: '',
        name: '',
        sub_district: {},
        district: {},
        province: {},
        default: "false",
        status: '',
        created_at: '',
        updated_at: ''
    });

    useEffect(() => {
        loadAddress();
        setSelectedAddress(address[0]);
        loadCookies();
    }, []);

    useEffect(() => {
        if (address.length > 0) {
            setSelectedAddress(address[0]);
        }
    }, [address]);

    const loadAddress = async () => {
        const token = AuthProvider.getAccessToken();
        try {
            const res = await getAddressAll(token || '');
            const fetchedAddresses = res.data.data;
            setAddress(fetchedAddresses); // Set the address state with the fetched data
            if (fetchedAddresses.length > 0) {
                setSelectedAddress(fetchedAddresses[0]); // Set the selected address to the first one
            }
            console.log("Address:", fetchedAddresses);
            setIsLoad(false);
        } catch (error) {
            console.error("An error occurred while fetching data:", error);
        }
    }
    const loadCookies = async () => {
        const checkoutData = Cookies.get("checkoutData");
        console.log("Checkout Data:", checkoutData);
        if (!checkoutData) {
            router.push("/cart");
            return;
        }
        const productData: ProductModel[] = JSON.parse(checkoutData);
        setProducts(productData);
        console.log("Product Data:", productData);
    }

    const handleAddressChange = (e: ChangeEvent<HTMLSelectElement>) => {
        const selected = address.find((item) => item.id === parseInt(e.target.value));
        if (selected) {
            setSelectedAddress(selected);
        }
        console.log("Selected Address:", selectedAddress);
    }

    const handleOrder = async () => {
        const token = AuthProvider.getAccessToken();
        const orderData: OrderCreate = {
            address_id: selectedAddress.id,
            type:"qr",
            total: products.reduce((acc, product) => acc + product.price, 0),
            // product_id * count added to the list_product_id
            list_product_id: products.flatMap(product => 
                Array(product.count).fill(product.id)
            ),
            img: image
        }
        console.log("Order Data:", orderData);
        try {
            const res = await postOrder(orderData, token || '');
            console.log("Order Response:", res);
            if (res.status === 200) {
                setAlertTitle("Success");
                setAlertContent("Order created successfully");
                setAlertStatus("success");
                setAlertConfirmText("OK");
                setAlertOnConfirm(() => () => {
                    Cookies.remove("checkoutData");
                    router.push("/");
                });
                setOpenAlert(true);
            } else {
                setAlertTitle("Error");
                setAlertContent("Failed to create order");
                setAlertStatus("error");
                setAlertConfirmText("OK");
                setAlertOnConfirm(() => {
                    setOpenAlert(false);
                });
                setOpenAlert(true);
            }
        } catch (error) {
            console.error("An error occurred while creating order:", error);
            setAlertTitle("Error");
            setAlertContent("Failed to create order");
            setAlertStatus("error");
            setAlertConfirmText("OK");
            setAlertOnConfirm(() => {
                setOpenAlert(false);
            });
            setOpenAlert(true);
        }
    }



    return (
        <>
        {
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
                        <Card>
                        <CardHeader>
                            <CardTitle>Your Order</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <CheckOutTable></CheckOutTable>
                            
                        </CardContent>
                        <CardContent>
                            
                            <Label htmlFor="address">Select Address</Label>
                                    <select id="address"  onChange={handleAddressChange} className="mt-1 form-select block w-full">
                                        {address.map((item) => (
                                        <option key={item.id} value={item.id}>
                                                {item.address_line}
                                            </option>
                                        ))}
                                    </select>
                            
                                    <div className="flex flex-col p-4">
                                        <h3 className="text-lg font-semibold">{selectedAddress.name}</h3>
                                        <p className="mt-1">{selectedAddress.address_line}</p>
                                        <div className="flex justify-between mt-1">
                                            <span>{selectedAddress.sub_district.name_en}</span>
                                            <span>({selectedAddress.sub_district.name_th})</span>
                                        </div>
                                        <div className="flex justify-between mt-1">
                                            <span>{selectedAddress.district.name_en}</span>
                                            <span>({selectedAddress.district.name_th})</span>
                                        </div>
                                        <div className="flex justify-between mt-1">
                                            <span>{selectedAddress.province.name_en}</span>
                                            <span>({selectedAddress.province.name_th})</span>
                                        </div>
                                        <p className="mt-1">Postal Code: {selectedAddress.sub_district.post_code}</p>
                                        <p className="mt-1">Phone: {selectedAddress.phone}</p>
                                    </div>
                        </CardContent>
                        <CardContent>
                            <Button onClick={() => setShowQRModal(true)}>Checkout</Button>
                        </CardContent>
                    </Card>
                    {showQRModal && (
                        <div className="fixed inset-0 bg-black bg-opacity-70 flex justify-center items-center">
                        <div className=" p-4 rounded-lg shadow-lg text-center">
                            <h3 className="font-bold text-lg mb-4">Scan QR Code</h3>
                            {/* Include your QR code here. For illustration, I'll use an image placeholder. */}
                            <img src="/path-to-your-qr-code-image.png" alt="QR Code" className="mb-4" />
                            <div>
                            {/* <p className="text-sm text-gray-500 mb-2">Total</p> */}
                            <label htmlFor="file-upload" className="block text-sm font-medium text-gray-500">
                                Upload payment proof 
                            </label>
                            <input
                                id="file-upload"
                                name="file-upload"
                                type="file"
                                required
                                className="mt-1 block w-full border border-gray-300 rounded-md p-2 text-sm"
                                // Implement file handling logic
                                onChange={(e) => setImage(e.target.files ? e.target.files[0] : null)}
                            />
                            </div>
                            <div className="mt-4">
                            <Button onClick={() => setShowQRModal(false)}>Close</Button>
                            <Button onClick={handleOrder}>Submit</Button>
                            </div>
                        </div>
                        </div>
                    )}
                </div>
            )
        }
        

    </>
    )
}
