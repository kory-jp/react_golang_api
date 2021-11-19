import { Box, Divider, Flex, Heading } from "@chakra-ui/layout";
import { useCallback, VFC } from "react"
import { useHistory } from "react-router";
import { PrimaryButton } from "../atoms/button/PrimaryButton";
import { TodoCard } from "../organisms/todo/TodoCard";

export const Index: VFC = () => {
  const history = useHistory()
  const onClickNewTodo = useCallback(()=> history.push("/todo/new"),[history])

  return(
    <Flex align="center" justify="center" mt="16">
      <Box 
        bg="white" 
        p="4" 
        w="4xl"
        borderRadius="md"
        shadow="md"
      >
        <Flex justify="space-between">
          <Heading mb="4">
            鈴木さんのTodos
          </Heading>
          <Box>
            <PrimaryButton
              onClick={onClickNewTodo}
            >
              Todo追加
            </PrimaryButton>
          </Box>
        </Flex>
        <Divider/>
        {/* ---- */}
        <TodoCard />
      </Box>
    </Flex>
  )
}

export default Index;