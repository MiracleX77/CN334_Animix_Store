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
import { getUser, putUser} from "@/apis/services/userService"
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



export default function EditAccount() {
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

    const [idUser, setIdUser] = useState('');
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');

    //data product
    

    useEffect(() => {
        loadUser();
    }, []);


    const loadUser = async () => {
        const fetchedUser = await getUser(token||'');
        const data = fetchedUser.data.data;
        setIdUser(data.id);
        setFirstName(data.first_name);
        setLastName(data.last_name);
        setEmail(data.email);
    };

    const handleFirstNameInputChange = (value: string) => {
        setFirstName(value);
    };
    const handleLastNameInputChange = (value: string) => {
        setLastName(value);
    }
    const handleEmailInputChange = (value: string) => {
        setEmail(value);
    }

    const handleSubmit = async (event:any) => {
        event.preventDefault();
        const data = {
            first_name: firstName,
            last_name: lastName,
            email: email
        }
        console.log(data);
        const res = await putUser(data, token||'');
        console.log(res);
        if(res.status === 200){
            setAlertTitle("Success");
            setAlertContent("Account updated successfully");
            setAlertConfirmText("OK");
            setAlertStatus("success");
            setAlertOnConfirm(() => () => setOpenAlert(false));
            setOpenAlert(true);
        } else {
            setAlertTitle("Failed");
            setAlertContent("Failed to update account");
            setAlertConfirmText("OK");
            setAlertStatus("error");
            setOpenAlert(true);
        }

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
            <Card>
            <CardHeader>
                <CardTitle>Edit Account</CardTitle>
                <CardDescription>Edit the Account details</CardDescription>
            </CardHeader>
            <CardContent>
                <form onSubmit={handleSubmit}>
                    <div className="grid grid-rows-2 grid-flow-col gap-4">

                    
                    <div className="flex row-span-3 col-span-1 ">
                        <User2 size={100} />
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="firstname">FirstName</Label>
                        <Input id="firstname" placeholder="firstname" value={firstName} onChange={(e) => handleFirstNameInputChange(e.target.value)} className="mt-1" />
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="lastname">LastName</Label>
                        <Input id="lastname" placeholder="lastname"value={lastName} onChange={(e) => handleLastNameInputChange(e.target.value)} className="mt-1" />
                    </div>
                    <div className="flex flex-col col-span-1">
                        <Label htmlFor="email">Email</Label>
                        <Input id="email" placeholder="email" value={email} onChange={(e) => handleEmailInputChange(e.target.value)} className="mt-1" />
                    </div>
                </div>
                <div className="flex justify-end col-span-3 space-x-2 mt-4">
                    <Button onClick={()=>router.push('/')}  className="bg-sec ring-2 px-8 rounded-xl text-white">
                        Back
                    </Button>
                    <Button type="submit" className="bg-primary ring-2 px-8 rounded-xl text-white">
                        Save
                    </Button>
                </div>
                </form>
            </CardContent>
        </Card>
        </>
    )
}