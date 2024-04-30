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
import { getAddressAll, getProvince} from "@/apis/services/addressService"
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
import { Address } from "@/models/dto/address"
import AddressAdd from "./addressAdd"



export default function AddressAccount() {
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

    const [address, setAddress] = useState<Address[]>([]);

    const [selectedProvince, setSelectedProvince] = useState('');


    //data product
    

    useEffect(() => {
        loadAddress();
        loadProvince();
    }, []);


    // const loadAuthor = async () => {
    //     const fetchedAuthor = await getAuthor(token||'');
    //     setSelectedAuthor(fetchedAuthor.data.data[0].id);
    //     setAuthor(fetchedAuthor.data.data);
    // };
    // const loadCategory = async () => {
    //     const fetchedCategory = await getCategory(token||'');
    //     setSelectedCategory(fetchedCategory.data.data[0].id);
    //     setCategory(fetchedCategory.data.data);
    // };
    const loadProvince = async () => {
        const fetchedPublisher = await getProvince(token||'');
        console.log(fetchedPublisher);
        // setSelectedPublisher(fetchedPublisher.data.data[0].id);
        // setPublisher(fetchedPublisher.data.data);
    };

    const loadAddress = async () => {
        const fetchedUser = await getAddressAll(token||'');
        const data : Address[] = fetchedUser.data.data;
        setAddress(data);
        console.log(data);
    };

    // const handleFirstNameInputChange = (value: string) => {
    //     setFirstName(value);
    // };
    // const handleLastNameInputChange = (value: string) => {
    //     setLastName(value);
    // }
    // const handleEmailInputChange = (value: string) => {
    //     setEmail(value);
    // }

    // const handleSubmit = async (event:any) => {
    //     event.preventDefault();
    //     const data = {
    //         first_name: firstName,
    //         last_name: lastName,
    //         email: email
    //     }
    //     console.log(data);
    //     const res = await putUser(data, token||'');
    //     console.log(res);
    //     if(res.status === 200){
    //         setAlertTitle("Success");
    //         setAlertContent("Account updated successfully");
    //         setAlertConfirmText("OK");
    //         setAlertStatus("success");
    //         setAlertOnConfirm(() => () => setOpenAlert(false));
    //         setOpenAlert(true);
    //     } else {
    //         setAlertTitle("Failed");
    //         setAlertContent("Failed to update account");
    //         setAlertConfirmText("OK");
    //         setAlertStatus("error");
    //         setOpenAlert(true);
    //     }

    //}

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
            <Card>
            <CardHeader>

            </CardHeader>
            {address.map((address: Address) => (
                <CardContent key={address.id}>
                <div className="flex flex-col p-4">
                    <h3 className="text-lg font-semibold">{address.name}</h3>
                    <p>{address.address_line}</p>
                    <p>{address.sub_district.name_en} ({address.sub_district.name_th})</p>
                    <p>{address.district.name_en} ({address.district.name_th})</p>
                    <p>{address.province.name_en} ({address.province.name_th})</p>
                    <p>Postal Code: {address.sub_district.post_code}</p>
                    <p>Phone: {address.phone}</p>
                    <p>{address.status === 'active' ? 'Active' : 'Inactive'}</p>
                    {/* If the 'default' property is meant to be a boolean, convert the string to a boolean */}
                    <p>{address.default === 'true' ? 'Default Address' : ''}</p>
                </div>
                </CardContent>
            ))}
        </Card>
        </>
    )
}