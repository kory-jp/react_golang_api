import { Button } from "@chakra-ui/button";
import { ReactNode, VFC } from "react";

type Props = {
  children: ReactNode;
  onClick: ()=> void;
}

export const PrimaryButton: VFC<Props> = (props) => {
  const {children, onClick} = props;

  return(
  <Button 
    w="100%" 
    bg="telegram.400" 
    olor="white"
    onClick={onClick}
  >
    {children}  
  </Button>
  )
}