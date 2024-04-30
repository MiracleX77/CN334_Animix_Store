'use client'
import React, { useState, useEffect } from 'react';
import { DataTable } from './data-table';
import { getOrderByStatus } from '@/apis/services/orderServices';
import { AuthProvider } from "@/utils/clientAuthProvider"
import { Order ,columns} from './columns';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { useRouter } from "next/navigation"
import {
    Tabs,
    TabsContent,
    TabsList,
    TabsTrigger,
} from "@/components/ui/tabs"
import { set } from 'date-fns';


function OrderTable() {

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

    const [status, setStatus] = useState('Pending');
    const [orders, setOrders] = useState([]);

    useEffect(() => {
        loadOrders("Pending");
    }, []);

    const loadOrders = async (status : string) => {

        const fetchedOrders = await getOrderByStatus(status,token||'');
        // Attach edit and delete methods to each order
        const order = fetchedOrders.data.data
        order.map((order: Order) => {
            if (order.status === "Pending") {
                order.view = () => router.push(`/dashboard/order/view/${order.id}`);
            } 
            if (order.status === "Processing") {
                order.view = () => router.push(`/dashboard/order/processing/${order.id}`);
            }
            if (order.status === "Shipped") {
                order.view = () => router.push(`/dashboard/order/shipped/${order.id}`);
            }
            if (order.status === "Delivered") {
                order.view = () => router.push(`/dashboard/order/delivered/${order.id}`);
            }
            order.created_at = new Date(order.created_at).toLocaleString();
        });
        setOrders(order);
    };

    const handleChangeStatus = (status:string) => {
        console.log("Change status to:", status);
        setStatus(status);
        loadOrders(status);
    }

    return (
    <>
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
        <Tabs defaultValue="pending" className="w-full">
            <TabsList className="grid w-full grid-cols-5">
                <TabsTrigger onClick={(event) => handleChangeStatus("Pending")} value="pending">Pending</TabsTrigger>
                <TabsTrigger onClick={(event) => handleChangeStatus("Processing")} value="processing">Processing</TabsTrigger>
                <TabsTrigger onClick={(event) => handleChangeStatus("Shipped")} value="shipped">Shipped</TabsTrigger>
                <TabsTrigger onClick={(event) => handleChangeStatus("Delivered")} value="delivered">Delivered</TabsTrigger>
                <TabsTrigger onClick={(event) => handleChangeStatus("Cancelled")} value="cancelled">Cancelled</TabsTrigger>
            </TabsList>
            <TabsContent value="pending">
                <DataTable columns={columns} data={orders} />
            </TabsContent>
            <TabsContent value="processing">
                <DataTable columns={columns} data={orders} />
            </TabsContent>
            <TabsContent value="shipped">
                <DataTable columns={columns} data={orders} />
            </TabsContent>
            <TabsContent value="delivered">
                <DataTable columns={columns} data={orders} />
            </TabsContent>
            <TabsContent  value="cancelled">
                <DataTable columns={columns} data={orders} />
            </TabsContent>
        </Tabs>
    </>
    
  );
}

export default OrderTable;