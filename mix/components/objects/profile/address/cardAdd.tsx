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
import { getProvince,getDistrict,getSubDistrict,postAddress} from "@/apis/services/addressService"
import { ChangeEvent, useEffect, useState } from "react"
import { AuthProvider } from "@/utils/clientAuthProvider"
import { set } from "date-fns"
import { Value } from "@radix-ui/react-select"
import Image from "next/image"
import { Button } from '@/components/ui/button';
import AlertDialog from "@/components/interactive/layout/alertdialog"
import { useRouter } from "next/navigation"
import { AddressCreate } from "@/models/dto/address"


export default function CardAdd() {
    const token = AuthProvider.getAccessToken()
    const router = useRouter();

    const [province, setProvince] = useState([{id:"",name_th:"",name_en:""}])
    const [district, setDistrict] = useState([{id:0,name_th:"",name_en:"",province_id:0}])
    const [subDistrict, setSubDistrict] = useState([{id:"",name_th:""}])


    const [openAlert, setOpenAlert] = useState(false);
    const [alertTitle, setAlertTitle] = useState('');
    const [alertContent, setAlertContent] = useState('');
    const [alertConfirmText, setAlertConfirmText] = useState('');
    const [alertOnConfirm, setAlertOnConfirm] = useState(() => () => console.log("default ooops"));
    const [alertOnCancel, setAlertOnCancel] = useState(() => () => console.log("default ooops"));
    const [alertStatus, setAlertStatus] = useState('');
    const [alertCancelBottom, setAlertCancelBottom] = useState(false);

    const [selectedSubDistrict, setSelectedSubDistrict] = useState(0);
    const [selectedDistrict, setSelectedDistrict] = useState(0);
    const [selectedProvince, setSelectedProvince] = useState(0);

    const [address, setAddress] = useState<AddressCreate>({
        address_line: "",
        province_id: 0,
        district_id: 0,
        sub_district_id: 0,
        phone: "",
        name: "",
        default:"false"
    });

    
    
    useEffect(() => {
        loadProvice();
        loadDistrict('1');
        loadSubDistrict('1001');
    }, []);


    const loadProvice = async () => {
        const fetchedProvince = await getProvince(token||'');
        setSelectedProvince(fetchedProvince.data.data[0].id);
        setProvince(fetchedProvince.data.data);
    };
    const loadDistrict = async (id:string) => {
        const fetchedDistrict = await getDistrict(token||'',parseInt(id));
        setSelectedDistrict(fetchedDistrict.data.data[0].id);
        setDistrict(fetchedDistrict.data.data);
    };
    const loadSubDistrict = async (id:string) => {
        const fetchedSubDistrict = await getSubDistrict(token||'',parseInt(id));
        setSelectedSubDistrict(fetchedSubDistrict.data.data[0].id);
        setSubDistrict(fetchedSubDistrict.data.data);
    };

    const handleNameInputChange = (value: string) => {
        const newAddress = { ...address };
        newAddress.name = value;
        setAddress(newAddress);
    };
    const handlePhoneInputChange = (value: string) => {
        const newAddress = { ...address };
        newAddress.phone = value;
        setAddress(newAddress);
    }
    
    const handleAddressLineInputChange = (value: string) => {
        const newAddress = { ...address };
        newAddress.address_line = value;
        setAddress(newAddress);
    }
    const handleSubDistrictChange = (event: { target: { value: any } }) => {
        const value = event.target.value;
        setSelectedSubDistrict(value);
    };

    const handleProvinceChange = (event: { target: { value: any } }) => {
        const value = event.target.value;
        loadDistrict(value);
        setSelectedProvince(value);
    };

    const handleDistrictChange = (event: { target: { value: any } }) => {
        const value = event.target.value;
        loadSubDistrict(value);
        setSelectedDistrict(value);
    };

    const handleSubmit = async (event:any) => {
        
        console.log(selectedProvince);
        console.log(selectedDistrict);
        console.log(selectedSubDistrict);
        event.preventDefault();

        var province_id = selectedProvince;
        var category_id = selectedDistrict;
        var subDistrict_id = selectedSubDistrict;


        const data = {
            name: address.name,
            phone: address.phone,
            address_line: address.address_line,
            province_id: province_id,
            district_id: category_id,
            sub_district_id: subDistrict_id,
            default: "false"
        }

        postAddress(data,token||'').then((res) => {
            if (res) {
                if (res.status === 200) {
                    console.log("success");
                    setAlertTitle("Address Added");
                    setAlertContent("Address added successfully");
                    setAlertConfirmText("OK");
                    setAlertStatus("success");
                    setAlertCancelBottom(false);
                    setAlertOnConfirm(() => () => {
                        setOpenAlert(false);
                        router.push('/dashboard/product');
                    });
                    setOpenAlert(true);
                } else {
                    console.log("error");
                    setAlertTitle("Add Address");
                    setAlertContent("Address addition failed");
                    setAlertConfirmText("OK");
                    setAlertStatus("error");
                    setAlertCancelBottom(false);
                    setAlertOnConfirm(() => () => setOpenAlert(false));
                    setOpenAlert(true);
                }

            } else{
                console.log('error');
            }
        })


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
                <CardTitle>Add Address</CardTitle>
                <CardDescription>Edit the product details</CardDescription>
            </CardHeader>
            <CardContent>
        <form onSubmit={handleSubmit} className="grid grid-cols-3 gap-4">

            <div className="flex flex-col col-span-1">
                <Label htmlFor="name">Name</Label>
                <Input id="name" placeholder="Name..." required onChange={(e) => handleNameInputChange(e.target.value)} className="mt-1" />
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="phone">Phone</Label>
                <Input id="phone" placeholder="Price..."  required onChange={(e) => handlePhoneInputChange(e.target.value)} className="mt-1" />
            </div>
            <div className="flex flex-col col-span-1">
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="province">Province</Label>
                    <select id="province" onChange={handleProvinceChange} className="mt-1 form-select block w-full">
                        {province.map((item) => (
                        <option key={item.id} value={item.id}>
                                {item.name_th}
                            </option>
                        ))}
                </select>
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="district">SubDistrict</Label>
                    <select id="district"  onChange={handleDistrictChange} className="mt-1 form-select block w-full">
                        {district.map((item) => (
                        <option key={item.id} value={item.id}>
                                {item.name_th}
                            </option>
                        ))}
                </select>
            </div>
            <div className="flex flex-col col-span-1">
                <Label htmlFor="sub_district">District</Label>
                    <select id="sub_district"  onChange={handleSubDistrictChange} className="mt-1 form-select block w-full">
                        {subDistrict.map((item) => (
                        <option key={item.id} value={item.id}>
                                {item.name_th}
                            </option>
                        ))}
                </select>
            </div>
            
            {/* ...more input fields... */}
            <div className="flex flex-col col-span-3">
                <Label htmlFor="description">Address Line</Label>
                <textarea id="description" onChange={(e) => handleAddressLineInputChange(e.target.value)} className="mt-1 h-24"></textarea>
            </div>
            <div className="col-span-3 flex justify-end space-x-2 mt-4">
                <Button onClick={()=>router.push('/dashboard/product')} className="bg-secondary ring-2 px-8 rounded-xl text-white">
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