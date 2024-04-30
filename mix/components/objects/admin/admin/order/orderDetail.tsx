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
import { getOrderById ,putOrder} from "@/apis/services/orderServices"
import { ChangeEvent, useEffect, useState } from "react"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { ProductModel } from "@/models/dto/product"
import { set } from "date-fns"
import { Value } from "@radix-ui/react-select"
import Image from "next/image"
import { Button } from '@/components/ui/button';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { ProductUpdate } from "@/models/dto/product"
import { useRouter } from "next/navigation"
import { User, User2 } from "lucide-react"
import { OrderModel, OrderUpdate } from "@/models/dto/order"



export default function OrderDetail({ id }: { id: string}) {
    const token = AuthProvider.getAccessToken()
    const router = useRouter();

    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertOnCancel, setAlertOnCancel] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    const [alertCancelBottom, setAlertCancelBottom] = useState(false);

    const [isLoad, setIsLoad] = useState(false);

    const [order, setOrder] = useState<OrderModel>({
        id: 0,
        user_id: 0,
        delivery: {
            id: 0,
            address: {
                id: 0,
                user_id: 0,
                address_line: '',
                phone: '',
                name: '',
                sub_district: {
                    id: 0,
                    name_th: '',
                    name_en: '',
                    post_code: '',
                },
                district: {
                    id: 0,
                    name_th: '',
                    name_en: '',
                },
                province: {
                    id: 0,
                    name_th: '',
                    name_en: '',
                },
                default: '',
                status: '',
                created_at: '',
                updated_at: '',
            },
            cost: 0,
            type: '',
            tracking_number: '',
            status: '',
            created_at: '',
            updated_at: '',
        },
        payment: {
            id: 0,
            type: '',
            proof_payment: '',
            status: '',
            created_at: '',
            updated_at: '',
            total: 0,
        },

        total_price: 0,
        status: '',
        created_at: '',
        updated_at: '',
    });

    //data product
    
    useEffect(() => {
        loadOrder();
    }, []);

    const loadOrder = async () => {
        const data = await getOrderById(id,token||'');
        setOrder(data.data.data);
        setIsLoad(false);
        
    };

    const updateStatus = () => {
        setAlertTitle('Update Status');
        setAlertContent('Are you sure you want to update this order status?');
        setAlertStatus('info');
        setAlertConfirmText('Yes');
        setAlertCancelBottom(true);
        setAlertOnConfirm(() => () => updateOrderStatus());
        setAlertOnCancel(() => () => setOpenAlert(false));
        setOpenAlert(true);
    }

    const updateOrderStatus = async () => {
        const data: OrderUpdate = {
            status: 'Processing',
        }
        const res = await putOrder(id,data,token||'');
        if (res.status === 200) {
            setAlertTitle('Success');
            setAlertContent('Order status has been updated');
            setAlertStatus('success');
            setAlertConfirmText('OK');
            setAlertOnConfirm(() => () => router.back());
            setOpenAlert(true);

        } else {
            setAlertTitle('Failed');
            setAlertContent('Order status update failed');
            setAlertStatus('error');
            setAlertConfirmText('OK');
            setAlertOnConfirm(() =>() => setOpenAlert(false));
            setOpenAlert(true);
        }
    }


    console.log(order);

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
                    <Card>
                    <CardHeader>
                
                    </CardHeader>
                    <CardContent>
                            <h3 className="text-md font-semibold">Delivery Details</h3>
                            <p><strong>Address:</strong> {order.delivery.address.address_line}</p>
                            <p><strong>Phone:</strong> {order.delivery.address.phone}</p>
                            <p><strong>Name:</strong> {order.delivery.address.name}</p>
                            <p><strong>District:</strong> {order.delivery.address.district.name_en} ({order.delivery.address.district.name_th})</p>
                            <p><strong>Province:</strong> {order.delivery.address.province.name_en} ({order.delivery.address.province.name_th})</p>
                            <p><strong>Postal Code:</strong> {order.delivery.address.sub_district.post_code}</p>
                            <p><strong>Status:</strong> {order.delivery.status}</p>
                            <p><strong>Tracking Number:</strong> {order.delivery.tracking_number || "N/A"}</p>

                            <h3 className="text-md font-semibold mt-4">Payment Information</h3>
                            <p><strong>Type:</strong> {order.payment.type}</p>
                            <p><strong>Total Paid:</strong> ${order.payment.total}</p>
                            <p><strong>Status:</strong> {order.payment.status}</p>
                            <p><strong>Proof of Payment:</strong> <a href={order.payment.proof_payment} target="_blank" rel="noopener noreferrer">View Payment</a></p>

                            <h3 className="text-md font-semibold mt-4">Order Summary</h3>
                            <p><strong>Total Price:</strong> ${order.total_price}</p>
                            <p><strong>Order Status:</strong> {order.status}</p>
                        </CardContent>
                        <CardContent>
                        <div className="flex justify-between p-4">
                            <button
                                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                                onClick={() => router.back()}
                            >
                                Back
                            </button>
                            <button
                                className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
                                onClick={updateStatus}
                            >
                                Update Status
                            </button>
                        </div>
                        </CardContent>

                        </Card>
                        </div>
            )
        }
            
        </>
    )
}