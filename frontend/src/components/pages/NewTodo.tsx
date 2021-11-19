import { Box, Flex } from "@chakra-ui/layout";
import { Textarea } from "@chakra-ui/textarea";
import { useCallback, VFC } from "react"
import { useHistory } from "react-router";
import { PrimaryButton } from "../atoms/button/PrimaryButton";

export const NewTodo: VFC = () => {
  const history = useHistory()
  const onClickIndex = useCallback(()=> history.push("/todo"),[history])
  return(
    <Flex align="center" justify="center" mt="16">
    <Box 
      bg="white" 
      p="4" 
      w="4xl"
      borderRadius="md"
      shadow="md"
    >
      <Textarea
        placeholder="Todoを追加"
        mb="4"
      />
      <PrimaryButton
        onClick={onClickIndex}
      >
        Todo追加
      </PrimaryButton>
    </Box>
  </Flex>
  )
}

export default NewTodo;