import * as a from "@/components/ui/alert-dialog"
import Middle from "@/components/layouts/Middle";
import { BadgeAlert, BadgeCheck, BadgeInfo, CheckIcon, CrossIcon } from "lucide-react";

import { useState } from "react";
type Props = {
    open: boolean;
    setOpen?: (open: boolean) => void;
    title: string;
    content?: string;
    status: string;
    onConfirm?: () => void;
    onCanceled?: () => void;
    confirmText: string;
    cancelBottom: boolean;
}

export default function AlertDialog(props: Props) {    
    return (
        <>
            <a.AlertDialog open={props.open} >
                <a.AlertDialogTrigger></a.AlertDialogTrigger>
                <a.AlertDialogContent>
                        <Middle X className="w-full h-full">
                        {
                            props.status === "success" ? (
                                <BadgeCheck className="w-10 h-10 text-Success" />
                            ) : props.status === "error" ? (
                                <BadgeAlert className="w-10 h-10 text-Danger" />
                            ) : props.status === "info" ? (
                                <BadgeInfo className="w-10 h-10 text-card-foreground" />
                            ) : props.status === "warning" ? (
                                <BadgeInfo className="w-10 h-10 text-Warning" />
                            ): (
                                <></>
                            )
                        }
                        </Middle>
                        <a.AlertDialogHeader>
                            <Middle X className="w-full h-full">
                                {props.title}
                            </Middle>
                        </a.AlertDialogHeader>
                        <a.AlertDialogDescription >
                            <Middle X className="w-full h-full">
                            {props.content}
                            </Middle>
                        </a.AlertDialogDescription>
                    <a.AlertDialogFooter>
                        <Middle X className="w-full h-full">
                        {
                        props.cancelBottom && <a.AlertDialogCancel 
                        className="bg-mauve4 hover:bg-mauve5 focus:shadow-mauve7 inline-flex h-[35px] items-center justify-center rounded-[4px] px-[15px] font-medium leading-none outline-none focus:shadow-[0_0_0_2px] mr-2" onClick={props.onCanceled}>
                        Cancel
                        </a.AlertDialogCancel>
                        }
                        {
                            props.confirmText && <a.AlertDialogAction className=" inline-flex h-[35px] items-center justify-center rounded-[4px] px-[15px] font-medium leading-none outline-none focus:shadow-[0_0_0_2px] " onClick={props.onConfirm}>{props.confirmText}</a.AlertDialogAction>
                        }
                        </Middle>
                    </a.AlertDialogFooter>
                    </a.AlertDialogContent>
            </a.AlertDialog>
        </>
    )
}