import Middle from "./Middle";

type Props = {
    children: React.ReactNode;
    isOpen: boolean;
};

const Modal = ({ children, isOpen }: Props) => {
    return (
        <>
        {
            isOpen && (
                <Middle X Y className="fixed top-0 left-0 w-screen h-screen bg-white bg-opacity-50 backdrop-blur-xl z-[999]">
                    {children}
                </Middle>
            )

        }
        </>
    );
}

export default Modal;
