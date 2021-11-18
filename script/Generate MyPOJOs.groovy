import com.intellij.database.model.DasTable
import com.intellij.database.util.Case
import com.intellij.database.util.DasUtil

/*
 * Available context bindings:
 *   SELECTION   Iterable<DasObject>
 *   PROJECT     project
 *   FILES       files helper
 */

packageName = ""
typeMapping = [
        (~/(?i)tinyint/)                            : "int8",
        (~/(?i)smallint/)                           : "int16",
        (~/(?i)int/)                                : "int",
        (~/(?i)bigint/)                             : "int64",
        (~/(?i)bool/)                               : "bool",
        (~/(?i)float|double|decimal|real/)          : "float64",
        (~/(?i)datetime|date|time|timestamp/)       : "*time.Time",
        (~/(?i)/)                                   : "string"
]

FILES.chooseDirectoryAndSave("Choose directory", "Choose where to store generated files") { dir ->
    SELECTION.filter { it instanceof DasTable }.each { generate(it, dir) }
}

def generate(table, dir) {
    def className = javaName(table.getName(), true)
    def fileName = table.getName()
    def fields = calcFields(table)
    packageName = getPackageName(dir)
    PrintWriter printWriter = new PrintWriter(new OutputStreamWriter(new FileOutputStream(new File(dir, fileName + ".go")), "UTF-8"))
    printWriter.withPrintWriter { out -> generate(out, className, fields) }
}

// 获取包所在文件夹路径
def getPackageName(dir) {
    return dir.toString().replaceAll("\\\\", ".").replaceAll("/", ".").replaceAll("^([^s]*)\\.", "") + ""
}

def generate(out, className, fields) {
    out.println "package $packageName"
    out.println ""

    Set types = new HashSet()

    fields.each() {
        types.add(it.type)
    }

    if (types.contains("*time.Time")) {
        out.println "import \"time\""
    }

    out.println ""
    out.println "type $className struct {"
    fields.each() {
        // 输出注释
        // 输出成员变量
        String comment = it.comment != null ? it.comment.toString() : ""
        out.println "\t${it.name}\t${it.type} // "+"${comment}"
    }
    out.println "}"
}

def calcFields(table) {
    DasUtil.getColumns(table).reduce([]) { fields, col ->
        def spec = Case.LOWER.apply(col.getDataType().getSpecification())
        def typeStr = typeMapping.find { p, t -> p.matcher(spec).find() }.value
        fields += [[
                           name : javaName(col.getName(), false),
                           type : typeStr,
                           comment: col.getComment(),
                           annos: ""]]
    }
}

def javaName(str, capitalize) {
    def s = com.intellij.psi.codeStyle.NameUtil.splitNameIntoWords(str)
            .collect { Case.LOWER.apply(it).capitalize() }
            .join("")
            .replaceAll(/[^\p{javaJavaIdentifierPart}[_]]/, "_")
    s
}
